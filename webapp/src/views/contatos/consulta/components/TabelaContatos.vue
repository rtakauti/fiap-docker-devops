<template>
  <b-table
    :fields="fields"
    @row-clicked="onRowClicked"
    @sort-changed="onSort" />
</template>

<script>
  export default {
    name: 'TabelaContatos',
    props: ['contacts', 'sorting'],
    data,
    methods: {
      onRowClicked,
      onSort
    }
  };

  function data() {
    return {
      fields: [
        {
          key: 'name',
          label: 'nome',
          sortable: true
        },
        {
          key: 'email',
          label: 'E-mail',
          sortable: true
        },
        {
          key: 'birth_date',
          label: 'Data de Nascimento',
          sortable: true
        },
        {
          key: 'phone',
          label: 'Celular',
          sortable: true
        }
      ]
    };
  }

  function onRowClicked() {
    console.log(arguments)
  }

  function onSort() {
    console.log(arguments)
  }
</script>
