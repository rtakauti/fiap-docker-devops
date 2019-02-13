<template>
  <TituloLayout titulo="Agenda de Contatos">
    <b-row>
      <b-col>
        <FiltroConsultaContatos
          @filter="onFilter"
          @reset="onFilter(null)" />
      </b-col>
    </b-row>

    <hr />

    <b-row>
      <b-col>
        <TabelaContatos
          :sorting="sort"
          :contacts-page="contacts"
          @sort="onSort" />
        <b-pagination v-if="totalPages > 1" />
      </b-col>
    </b-row>
  </TituloLayout>
</template>

<script>
  import axios from 'axios';
  import FiltroConsultaContatos from './components/FiltroConsultaContatos.vue';
  import TabelaContatos from './components/TabelaContatos.vue';

  export default {
    name: 'ConsultaContatos',
    components: { FiltroConsultaContatos, TabelaContatos },
    data,
    methods: {
      onFilter,
      toPage,
      onSort,
      _getContactPage
    }
  };

  function data() {
    return {
      filter: null,
      contacts: null,
      page: 0,
      sort: null,
      limit: 10,
      totalCount: 0,
      totalPages: 0
    };
  }

  function onFilter(filter) {
    this.page = 0;
    this.sort = null;
    this.filter = filter;
    this._getContactPage();
  }

  function toPage(page) {
    this.page = page;
    this._getContactPage();
  }

  function onSort(sort) {
    this.sort = sort;
    this._getContactPage();
  }

  function _getContactPage() {
    const params = {
      _page: this.page,
      _limit: this.limit
    };

    if (this.sort) {
      params._sort = this.sort.by;
      params._order = this.sort.order;
    }

    if (this.filter != null) {
      Object.keys(this.filter).forEach(key => params[key] = this.filter[key]);
    }

    axios.get('/contacts', { params })
      .then(({ data, headers }) => {
        this.totalCount = headers['x-total-count'];
        this.totalPages = this.totalCount / this.limit;
        this.contacts = data;
      });
  }
</script>
