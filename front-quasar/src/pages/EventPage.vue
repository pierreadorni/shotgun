<template>
  <q-page>
    <div class="event-container">
      <div class="event-content">
        <div class="flex row items-center q-gutter-x-sm">
          <h5>{{ event.Title }}</h5>
          <q-btn
            round
            flat
            icon="edit"
            color="primary"
            @click="()=>goToEditEvent(event.ID)"
          />
        </div>
        <div class="event-description">
          {{ event.Description }}
        </div>
        <div class="row q-gutter-x-md">
          <q-circular-progress
            :min="0"
            :max="event.MaxSubscribers"
            :value="event.Subscribers.length"
            size="50px"
            :thickness="0.22"
            color="teal"
            track-color="grey-3"
            class="q-ma-md"
          />
          <div class="flex column items-center">
            <div class="text-h6">{{ event.Subscribers.length }}</div>
            <div class="text-subtitle1">Subscribers</div>
          </div>
        </div>
        <div class="flex row items-center q-gutter-x-sm">
          <q-icon name="event"/>
          <div class="text-subtitle1">
            {{ (new Date(event.StartTime)).toLocaleDateString() }} - {{
              (new Date(event.EndTime)).toLocaleDateString()
            }}
          </div>
        </div>
        <div class="event-sidebar">
        </div>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { Event } from 'src/types/event';

export default defineComponent({
  name: 'EventPage',
  setup() {
    const router = useRouter();
    const eventId = router.currentRoute.value.params.id;
    return {
      eventId,
    };
  },
  data() {
    return {
      event: {} as Event,
    };
  },
  methods: {
    fetchEvent() {
      axios
        .get(`http://localhost:3000/api/events/${this.eventId}`)
        .then((res) => {
          this.event = res.data as Event;
        });
    },
  },
  created() {
    this.fetchEvent();
  },
});
</script>

<style scoped>

.event-container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.event-content {
  width: 100%;
  padding-right: 40px;
}

.event-sidebar {
  width: 200px;
  border-left: 1px solid #e0e0e0;
}

.event-description {
  margin-top: 20px;
  color: #9e9e9e;
  text-align: justify;
}

</style>
