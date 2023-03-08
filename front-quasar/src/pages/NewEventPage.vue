<template>
  <q-page>
    <h5
      class="text-h5"
    >
      Create a new event
    </h5>
    <q-form
      class="q-gutter-y-md"
      @submit="createEvent"
    >
      <div class="q-gutter-x-md row items-start">
        <q-input
          label="Title *"
          filled
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
          v-model="title"
        />
        <q-input
          label="Location"
          filled
          v-model="location"
        />
        <q-input
          label="Start Date *"
          filled
          v-model="startDate"
          type="date"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please choose a Date']"
        />
        <q-input
          label="End Date *"
          filled
          v-model="endDate"
          type="date"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please choose a Date']"
        ></q-input>
        <q-input
          label="Max Subscribers"
          filled
          v-model="maxSubscribers"
          type="number"
          lazy-rules
          :rules="[ val => val && val >= 0 || 'Please type a number 0 or greater']"
        />
      </div>
      <q-input
        label="Description *"
        filled
        v-model="description"
        type="textarea"
        lazy-rules
        :rules="[ val => val && val.length > 0 || 'Please type something']"
      />
      <q-btn
        label="Create"
        color="primary"
        type="submit"
      />
    </q-form>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';
import { Event } from 'src/types/event';
import { Session } from 'src/types/session';
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar';

export default defineComponent({
  name: 'NewEventPage',
  setup() {
    const location = ref('');
    const title = ref('');
    const description = ref('');
    const startDate = ref('');
    const endDate = ref('');
    const maxSubscribers = ref(0);
    const userId = '';
    const $q = useQuasar();
    const router = useRouter();

    const createEvent = async () => {
      const evt: Event = {
        Title: title.value,
        Description: description.value,
        Location: location.value,
        StartTime: startDate.value,
        EndTime: endDate.value,
        Owner: userId,
        Subscribers: [],
        MaxSubscribers: Number(maxSubscribers.value),
      };
      await axios.post('http://localhost:3000/api/events', evt, {
        withCredentials: true,
      }).then((res) => {
        if (res.status === 201) {
          $q.notify({
            message: 'Event created successfully',
            color: 'positive',
            position: 'top',
          });
          router.push(`/events/${res.data.ID}`);
        }
      }).catch((err) => {
        $q.notify({
          message: `Error creating event: ${err.response.data}`,
          color: 'negative',
          position: 'top',
        });
      });
    };

    return {
      location,
      title,
      description,
      startDate,
      endDate,
      userId,
      maxSubscribers,
      createEvent,
    };
  },
  created() {
    axios.get('http://localhost:3000/session', {
      withCredentials: true,
    }).then((res) => {
      this.userId = (res.data as Session).uid;
    });
  },
});
</script>
