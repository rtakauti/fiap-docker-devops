<template>
  <div>
    <b-row>
      <b-col>
        <h1>{{ titulo }}</h1>
      </b-col>
    </b-row>
    <slot />
  </div>
</template>

<script>
  export default {
    name: 'TituloLayout',
    props: ['titulo']
  };
</script>
