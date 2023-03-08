<script lang="ts">
import Button from "@/components/Button.vue";
import Spinner from "@/components/Spinner.vue";
import {defineComponent} from "vue";

export default defineComponent({
  name: "LoginButton",
  components: {Spinner, Button},
  data() {
    return {
      loading: false,
      loggedIn: null,
    }
  },
  created() {
    this.$watch(
        () => this.$route.params,
        () => {
          this.fetchData()
        },
        { immediate: true }
    )
  },
  methods: {
    fetchData() {
      this.loading = true
      fetch("http://localhost:3000/session", {credentials: "include"}).then((response) => {
        this.loggedIn = response.ok
        this.loading = false
      }).catch((error) => {
        this.loading = false
        console.error(error)
      })
    },
  },
})
</script>

<template>
  <a v-if="!loading && !loggedIn" href="http://localhost:3000/login">
    <Button>connexion</Button>
  </a>
  <a v-if="!loading && loggedIn" href="http://localhost:3000/logout">
    <Button>déconnexion</Button>
  </a>
  <Button v-if="loading" disabled><Spinner/></Button>
</template>

<style scoped>

</style>