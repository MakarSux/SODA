<template>
  <header class="header">
    <div class="container">
      <div class="header__content">
        <div class="header__logo">
          <h1>SODA</h1>
        </div>
        <nav class="header__nav">
          <a href="#services">Услуги</a>
          <a href="#portfolio">Портфолио</a>
          <a href="#contacts">Контакты</a>
        </nav>
      </div>
    </div>
  </header>
</template>

<script setup>
// Component logic here
</script>

<style scoped lang="scss">
.header {
  padding: 1rem 0;
  background: #fff;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  position: fixed;
  width: 100%;
  top: 0;
  z-index: 1000;

  &__content {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  &__logo {
    h1 {
      font-size: 2rem;
      font-weight: bold;
      color: #333;
      margin: 0;
    }
  }

  &__nav {
    display: flex;
    gap: 2rem;

    a {
      text-decoration: none;
      color: #333;
      font-weight: 500;
      transition: color 0.3s ease;

      &:hover {
        color: #007bff;
      }
    }
  }
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

@media (max-width: 768px) {
  .header {
    &__nav {
      gap: 1rem;
      
      a {
        font-size: 0.9rem;
      }
    }
  }
}
</style>

