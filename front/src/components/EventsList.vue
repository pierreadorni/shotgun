<template>
  <div v-if="loading">
    loading...
  </div>
  <div v-else>
    <h2>Events: </h2>
    <EventCard :event="event" v-for="event in events" key="event.ID" />
  </div>
</template>

<script lang="ts">
import type { Event } from "@/types/event"
import EventCard from "@/components/EventCard.vue";
export default {
  name: "EventsList",
  components: {EventCard},
  data() {
    return {
      events: [] as Event[],
      loading: false,
    }
  },
  created() {
    this.fetchEvents()
  },
  methods: {
    fetchEvents() {
      this.loading = true
      fetch("http://localhost:3000/api/events").then((response) => {
        return response.json()
      }).then((events) => {
        this.events = events
        this.loading = false
      }).catch((error) => {
        this.loading = false
        console.error(error)
      })
    },
  },
}
</script>