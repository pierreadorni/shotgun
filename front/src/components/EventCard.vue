<template>
  <RouterLink :to="'/event/'+event.ID">
    <div class="event-card">
      <h2>{{ event.Title }}</h2>
      <p>{{ event.Description }}</p>
    </div>
  </RouterLink>
</template>

<script lang="ts">
import type Event from "@/types/event"

export default {
  name: "EventCard",
  props: {
    event: {
      type: Object as () => Event,
      required: true,
    },
  },
}
</script>

<style scoped>

.event-card {
  background-color: #4CAF50;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  transition: all 0.3s ease 0s;
  cursor: pointer;
}

</style>