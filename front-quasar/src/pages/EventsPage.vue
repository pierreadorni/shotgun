<template>
  <q-page>
    <q-list class="q-pa-md row items-start q-gutter-md">
      <q-card
        v-for="event in events"
        :key="event.ID"
        class="q-mb-bd my-card"
        flat
        bordered
        >
        <q-card-section class="bg-primary text-white">
          <div class="text-h6">{{ event.Title }}</div>
          <div class="text-subtitle">by {{ event.Owner }}</div>
        </q-card-section>
        <q-card-section class="text-justify">
          {{ truncate(event.Description, 200) }}
        </q-card-section>
        <q-card-actions class="absolute-bottom">
          <q-separator/>
          <q-btn flat @click="()=>goToEvent(event.ID)">View</q-btn>
        </q-card-actions>
      </q-card>
    </q-list>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Event } from 'src/types/event';
import axios from 'axios';

export default defineComponent({
  name: 'EventsPage.vue',
  data() {
    return {
      events: [] as Event[],
    };
  },
  async created() {
    const response = await axios.get('http://localhost:3000/api/events');
    this.events = response.data;
  },
  methods: {
    truncate(str: string, length: number) {
      return str.length > length ? `${str.substring(0, length)}...` : str;
    },
    goToEvent(eventId: number) {
      this.$router.push(`/events/${eventId}`);
    },
  },
});
</script>

<style scoped>
.my-card {
  width: 400px;
  height: 250px;
}

</style>
