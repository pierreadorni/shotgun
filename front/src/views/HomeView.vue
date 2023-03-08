<template>
  <main>
    <EventsList />
    <div>
      <Button>
        <v-icon name="md-add-circle"/> New Event
      </Button>
    </div>
  </main>
</template>

<script setup lang="ts">
import EventsList from "@/components/EventsList.vue";
import Button from "@/components/Button.vue";
</script>


<style scoped>

main {
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding: 20px;
}

</style>