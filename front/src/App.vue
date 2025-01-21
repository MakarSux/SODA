<template>
  
    <Header />
    
    <main>
      <section class="hero">
        <div class="container">
          
          <a href="#request-form" class="cta-button">Оставить заявку</a>
        </div>
      </section>

      <section id="services" class="services">
        <div class="container">
          <h2>Наши услуги</h2>
          <BannerService />
        </div>
      </section>

      <section id="request-form" class="request-section">
        <div class="container">
          <RequestForm />
        </div>
      </section>
    </main>

    <footer class="footer">
      <div class="container">
        <p>&copy; 2025 Веб-студия SODA. Все права защищены.</p>
      </div>
    </footer>
  
</template>

<script setup>
import Header from './components/Header.vue'
import BannerService from './components/BannerService.vue'
import RequestForm from './components/RequestForm.vue'
</script>

<style lang="scss">
// Reset and base styles
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Arial', sans-serif;
  line-height: 1.6;
  color: #333;
  background: #f8f9fa;
}

main{
  margin: 0 auto;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

// Hero section
.hero {
  width: 100%;
  min-height: 850px;

  background: url('../src/assets/background.png');
  background-size: cover;
  background-position: center;
  color: white;

  text-align: center;
  margin-top: 60px;

  h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
  }

  p {
    font-size: 1.2rem;
    margin-bottom: 2rem;
  }
}

.cta-button {
  display: inline-block;
  padding: 1rem 2rem;
  background: white;
  color: #007bff;
  text-decoration: none;
  border-radius: 4px;
  font-weight: bold;
  transition: transform 0.3s ease;

  &:hover {
    transform: translateY(-2px);
  }
}

// Services section
.services {
  padding: 4rem 0;
  background: white;

  h2 {
    text-align: center;
    margin-bottom: 3rem;
    font-size: 2.5rem;
  }
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.service-card {
  background: #f8f9fa;
  padding: 2rem;
  border-radius: 8px;
  text-align: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  transition: transform 0.3s ease;

  &:hover {
    transform: translateY(-5px);
  }

  i {
    font-size: 2.5rem;
    color: #007bff;
    margin-bottom: 1rem;
  }

  h3 {
    margin-bottom: 1rem;
  }
}

// Request form section
.request-section {
  padding: 4rem 0;
  background: #f8f9fa;
}

// Footer
.footer {
  background: #333;
  color: white;
  padding: 2rem 0;
  text-align: center;
}

// Responsive design
@media (max-width: 768px) {
  .hero {
    padding: 6rem 0 4rem;

    h1 {
      font-size: 2.5rem;
    }
  }

  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>
