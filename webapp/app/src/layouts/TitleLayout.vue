<template>
  <div>
    <b-row>
      <b-col>
        <h1>{{ title }}</h1>
      </b-col>
    </b-row>
    <slot />
  </div>
</template>

<script>
  export default {
    name: 'TitleLayout',
    props: ['title']
  };
</script>
