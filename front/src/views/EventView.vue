<template>
  <main>
    <div class="main-content">
      <h2>{{ event.Title }}</h2>
      <p>{{ event.Description }}</p>
      <span>created by {{ event.Owner }}</span>
    </div>
    <div class="side-bar">
      <div class="section">
        <span>created {{ event.CreatedAt.toLocaleDateString() }}</span>
      </div>
      <div class="section">
        <Button variant="danger" @click="deleteEvent" :disabled="session.uid !== event.Owner">Delete</Button>
      </div>
    </div>
  </main>
</template>


<script lang="ts">
import {defineComponent} from "vue";
import Button from "@/components/Button.vue";
import {getEvent} from "@/helpers/events";
import type {Event} from "@/types/event";
import type {Session} from "@/types/session";

export default defineComponent({
  name: "EventView",
  components: {Button},
  data() {
    return {
      event: {} as Event,
      loading: false,
      session: {} as Session
    }
  },
  created() {
    this.fetchEvent()
    this.fetchSession()
  },
  methods: {
    fetchEvent() {
      this.loading = true
      getEvent(this.$route.params.id as string).then((event) => {
        console.log(event)
        this.event = event
        this.loading = false
      }).catch((error) => {
        console.error(error)
        this.loading = false
      })
    },
    deleteEvent() {
      fetch("http://localhost:3000/api/events/" + this.$route.params.id, {
        method: "DELETE",
        credentials: "include",
      }).then((response) => {
        return response.json()
      }).then((event) => {
        this.$router.push("/")
      }).catch((error) => {
        console.error(error)
      })
    },
    fetchSession() {
      fetch("http://localhost:3000/session", {credentials: "include"})
          .then(res => res.json())
          .then((session) => {
            this.session = session
          }).catch((error) => {
        console.error(error)
      })
    }
  },
})
</script>

<style scoped>

main {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 20px;
  height: 100%;
}

.side-bar {
  width: 300px;
  margin-left: 20px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  border-left: 1px solid #ddd;
  padding-left: 20px;
}

.section {
  margin-bottom: 20px;
}

.section:not(:last-child) {
  border-bottom: 1px solid #ddd;
  padding-bottom: 20px;
}

</style>