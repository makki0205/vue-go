<template>
  <section id="sidebar">
    <h1>recipo</h1>
    <ul @click="select">
      <li v-for="item in items" :class="{ active: item.isActive }">{{item.name}}</li>
    </ul>
    <div class="container">
      <div class="profile">
        <div class="img">
          <img src="https://github.com/konojunya.png" alt="icon">
        </div>
        <p class="username">botti</p>
      </div>
      <p class="user-info">ユーザー情報</p>
      <p class="logout">ログアウト</p>
      <small>&copy; copyright</small>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      items: [
        { name: "全てのレシピ", isActive: true },
        { name: "ブックマーク", isActive: false },
        { name: "小腹が空いた時", isActive: false },
        { name: "カットの仕方", isActive: false },
        { name: "お手軽夜食", isActive: false }
      ]
    }
  },
  components: {
  },
  methods: {
    select(e) {
      for(let item of this.items) {
        item.isActive = e.target.textContent === item.name
      }
    }
  }
}
</script>

<style lang="scss" scoped>
#sidebar {
  background-color: #272727;
  width: 15%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  box-shadow: 5px 0px 15px black;
  h1 {
    font-size: 1.2rem;
    font-weight: 100;
    letter-spacing: 1.5px;
    color: white;
    padding: 15px 30px;
  }
  ul {
    width: 100%;
    color: white;
    font-size: 13px;
    cursor: pointer;
    li {
      padding: 15px 30px;
      list-style: none;
      &.active {
        background-color: rgba(255,255,255,0.15);
      }
    }
  }
  .container {
    margin-top: auto;
    color: white;
    font-size: 13px;
    padding: 0 30px;
    .profile {
      width: 100%;
      .img {
        width: 50%;
        overflow: hidden;
        border-radius: 50%;
        border: 1px solid #979797;
        img {
          display: block;
          width: 100%;
        }
      }
      .username {
        margin: 20px 0;
      }
    }
    .user-info {
      cursor: pointer;
    }
    .logout {
      cursor: pointer;
      margin: 5px 0;
    }
    small {
      color: gray;
    }
  }
}
</style>

