package routes

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gopkg.in/cas.v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var store *session.Store

type CasAuthenticationFailure struct {
	Code    string `xml:"code,attr"`
	Message string `xml:",chardata"`
}

type CasUserAttributes struct {
	Uid            string `xml:"uid" json:"uid"`
	Mail           string `xml:"mail" json:"mail"`
	AccountProfile string `xml:"accountProfile" json:"account_profile"`
	DisplayName    string `xml:"displayName" json:"display_name"`
	OU             string `xml:"ou" json:"ou"`
	GivenName      string `xml:"givenName" json:"given_name"`
	CN             string `xml:"cn" json:"cn"`
	SN             string `xml:"sn" json:"sn"`
}

type CasAuthenticationSuccess struct {
	User       string            `xml:"user"`
	Attributes CasUserAttributes `xml:"attributes"`
}
type CasServiceResponse struct {
	AuthenticationFailure CasAuthenticationFailure `xml:"authenticationFailure"`
	AuthenticationSuccess CasAuthenticationSuccess `xml:"authenticationSuccess"`
}

func logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// delete the account id from the session
	err = sess.Destroy()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// logout from CAS
	casUrl, err := url.Parse("https://cas.utc.fr/cas")
	if err != nil {
		return err
	}

	casClient := cas.NewClient(&cas.Options{
		URL: casUrl,
	})
	if casClient == nil {
		err := "cas: redirect to cas failed as no client associated with request"
		return fiber.NewError(http.StatusInternalServerError, err)
	}
	redirectUrl, err := url.Parse("http://localhost:3000/login")
	if err != nil {
		return err
	}
	u, err := casClient.LogoutUrlForRequest(&http.Request{URL: redirectUrl, Host: "localhost:3000"})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(u, http.StatusFound)

}

func loginCas(c *fiber.Ctx) error {
	var req http.Request
	err := fasthttpadaptor.ConvertRequest(c.Context(), &req, false)
	if err != nil {
		return err
	}
	casUrl, err := url.Parse("http://cas.utc.fr/cas")
	if err != nil {
		return err
	}

	if !cas.IsAuthenticated(&req) {
		casClient := cas.NewClient(&cas.Options{
			URL: casUrl,
		})
		if casClient == nil {
			err := "cas: redirect to cas failed as no client associated with request"
			return fiber.NewError(http.StatusInternalServerError, err)
		}

		//c.RedirectToLogin(w, r)
		u, err := casClient.LoginUrlForRequest(&req)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(u, http.StatusFound)
	}

	return c.Next()
}

func validateTicket(c *fiber.Ctx) error {
	// if there is a ?ticket= in the url, validate the ticket
	ticket := c.Query("ticket")
	if ticket != "" {
		// check whether the ticket is valid
		casUrl, err := url.Parse("https://cas.utc.fr/cas")
		if err != nil {
			return err
		}
		casClient := cas.NewClient(&cas.Options{
			URL: casUrl,
		})
		if casClient == nil {
			err := "cas: redirect to cas failed as no client associated with request"
			return fiber.NewError(http.StatusInternalServerError, err)
		}
		redirectUrl, err := url.Parse("http://localhost:3000/login")
		if err != nil {
			return err
		}
		// get the account id from the ticket
		validateUrl, err := casClient.ServiceValidateUrlForRequest(ticket, &http.Request{URL: redirectUrl, Host: "localhost:3000"})
		if err != nil {
			return err
		}
		// get the url
		resp, err := http.Get(validateUrl)
		if err != nil {
			return err
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// parse the response
		var casResponse CasServiceResponse
		err = xml.Unmarshal(bodyBytes, &casResponse)
		if err != nil {
			return err
		}

		sess, err := store.Get(c)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		// set values in the session
		attrs, err := json.Marshal(casResponse.AuthenticationSuccess.Attributes)
		if err != nil {
			return err
		}
		sess.Set("attributes", string(attrs))
		err = sess.Save()
		// ticket validated
		return c.Redirect(os.Getenv("FRONT_ADDR"), http.StatusFound)
	}
	return c.Next()
}

func visibleByAccountId(c *fiber.Ctx) error {
	// get the account id from the sess
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Failed to get session: "+err.Error())
	}
	accountId := sess.Get("id")
	if accountId == nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "You are not logged in")
	}

	// get the account id from the url
	id := c.Params("id")
	// check if the account id from the request context matches the account id from the url
	if strconv.Itoa(int(accountId.(uint))) != id {
		return fiber.NewError(fiber.ErrForbidden.Code, "Your account id ("+strconv.Itoa(int(accountId.(uint)))+") does not match the account id in the url")
	}
	return c.Next()
}
