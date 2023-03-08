<template>
<q-page class="flex flex-center q-gutter-y-md" style="flex-direction: column">
  <span v-if="session.uid !== undefined">connecté en tant que {{ session.display_name }}</span>
  <div class="q-gutter-x-md">
    <q-btn color="primary" @click="redirectToLogin" :disable="session.uid !== undefined">
      Connexion
    </q-btn>
    <q-btn color="secondary" @click="redirectToLogout" :disable="session.uid === undefined">
      Déconnexion
    </q-btn>
  </div>
</q-page>
</template>

<script lang="ts">
import { Session } from 'src/types/session';
import getSession from 'src/helpers/session';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'LoginPage',
  data: () => ({
    session: {} as Session,
  }),
  methods: {
    redirectToLogin() {
      window.location.href = 'http://localhost:3000/login';
    },
    redirectToLogout() {
      window.location.href = 'http://localhost:3000/logout';
    },
  },
  async created() {
    this.session = await getSession();
    console.log(this.session);
  },
});
</script>

<style scoped>

</style>
