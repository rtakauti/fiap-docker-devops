<template>
  <b-table
    :sort-by="sortBy"
    :sort-desc="sortDesc"
    hover
    :fields="fields"
    :items="contacts"
    @row-clicked="onRowClicked"
    @sort-changed="onSort" />
</template>

<script>
  export default {
    name: 'ContactsSearchTable',
    props: ['contacts', 'sorting'],
    data,
    computed: {
      sortBy,
      sortDesc
    },
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
          label: 'Nome',
          sortable: true
        },
        {
          key: 'email',
          label: 'E-mail',
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

  function sortBy() {
    return this.sorting != null ?  this.sorting.by : null;
  }

  function sortDesc() {
    return this.sorting != null ?  this.sorting.order === 'DESC' : false;
  }

  function onRowClicked({ id }) {
    this.$router.push({
      name: 'contacts-edit',
      params: {
        id
      }
    });
  }

  function onSort({ sortBy, sortDesc }) {
    this.$emit('sort', {
      by: sortBy,
      order: sortDesc ? 'DESC' : 'ASC'
    });
  }
</script>
