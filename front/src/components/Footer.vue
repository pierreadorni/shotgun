<template>
  <span v-if="!loading && id">
    logged in as {{ id }}
  </span>
  <span v-if="!loading && !id">
    not logged in
  </span>
  <span v-if="loading">
    loading...
  </span>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "Footer",
  data() {
    return {
      loading: false,
      id: "",
    }
  },
  created() {
    this.$watch(
        () => this.$route.params,
        () => {
          this.fetchData()
        },
        {immediate: true}
    )
  },
  methods: {
    fetchData() {
      this.loading = true
      fetch("http://localhost:3000/session", {credentials: "include"})
          .then(res => res.json())
          .then((session) => {
            this.id = session.uid
            this.loading = false
          }).catch((error) => {
        this.loading = false
        console.error(error)
      })
    },
  },
})
</script>

<style scoped>

span {
  opacity: 0.5;
}

</style>