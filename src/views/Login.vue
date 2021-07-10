<template>
  <div style="margin-top:20px">
    <h3>
      {{ judulHalaman }}
    </h3>

    <form>
      <input type="text" name="" id="" placeholder="User/email" />
      <input type="password" name="" id="" placeholder="password" />
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      judulHalaman: 'Login',
    }
  },

  methods: {
    getJudul() {
      this.judulHalaman = 'JUDUL BARU'
    },
  },
}
</script>
