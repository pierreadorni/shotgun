<template>
  <button :class="variant" :disabled="disabled">
    <slot></slot>
  </button>
</template>

<script lang="ts">
import {defineComponent} from "vue";

type ButtonVariant = "primary" | "danger";

export default defineComponent({
  name: "Button",
  props: {
    variant: {
      type: String as () => ButtonVariant,
      default: "primary",
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  },
})
</script>

<style scoped>
  button {
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

  .primary {
    background-color: #4CAF50;
  }

  .danger {
    background-color: #f44336;
  }

  button.primary:hover {
    background-color: #3e8e41;
  }

  button.danger:hover {
    background-color: #da190b;
  }

  button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }

  button:disabled:hover {
    background-color: #ccc;
  }
</style>