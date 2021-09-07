<template>
  <div id="app">
    <order-summary></order-summary>
  </div>
</template>

<script>

import OrderSummary from "./components/OrderSummary";

export default {
  name: 'App',
  components: {
    OrderSummary,
  },
}
</script>

<style>
#app {
  font-family: 'Red Hat Display', sans-serif;
}

</style>
