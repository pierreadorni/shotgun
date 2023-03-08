<template>
  <q-layout view="hHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Quasar App
        </q-toolbar-title>

        <div>Quasar v{{ $q.version }}</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
<!--        <EssentialLink-->
<!--          v-for="link in essentialLinks"-->
<!--          :key="link.title"-->
<!--          v-bind="link"-->
<!--        />-->
        <EssentialLink
          title="List Events"
          caption="A list of all the events created by anyone"
          icon="list"
          link="/events"
        />
        <EssentialLink
          title="My Events"
          caption="A list of all the events you created"
          icon="event"
          link="/my-events"
        />
        <q-separator inset />
        <EssentialLink
          title="New Event"
          icon="add"
          link="/events/new"
        />
        <q-separator inset />
        <EssentialLink
          title="Login"
          icon="login"
          link="/login"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';

const linksList = [
  {
    title: 'List Events',
    caption: 'A list of all the events created by anyone',
    icon: 'list',
    link: '/events',
  },
  {
    title: 'My Events',
    caption: 'A list of all the events you created',
    icon: 'event',
    link: '/my-events',
  },
];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink,
  },

  setup() {
    const leftDrawerOpen = ref(false);

    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
    };
  },
});
</script>

<style>
.q-page{
  padding: 50px;
}
</style>
