<template>
  <div class="general-view">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-content">
        <h1>SODA Design</h1>
        <p class="subtitle">Ваш надежный партнер в мире дизайна</p>
        <div class="cta-buttons">
          <router-link to="/register" class="cta-button primary">Начать работу</router-link>
          <router-link to="/about_us" class="cta-button secondary">Узнать больше</router-link>
        </div>
      </div>
    </section>

    <!-- Services Section -->
    <section class="services">
      <h2>Наши услуги</h2>
      <div class="services-grid">
        <div class="service-card">
          <i class="fas fa-paint-brush"></i>
          <h3>Графический дизайн</h3>
          <p>Создание логотипов, брендбуков и фирменного стиля для вашего бизнеса</p>
        </div>
        <div class="service-card">
          <i class="fas fa-desktop"></i>
          <h3>Веб-дизайн</h3>
          <p>Разработка современных и удобных интерфейсов для веб-сайтов</p>
        </div>
        <div class="service-card">
          <i class="fas fa-mobile-alt"></i>
          <h3>Мобильные приложения</h3>
          <p>Дизайн пользовательских интерфейсов для мобильных приложений</p>
        </div>
        <div class="service-card">
          <i class="fas fa-video"></i>
          <h3>Моушн-дизайн</h3>
          <p>Создание анимаций и видеоконтента для вашего бренда</p>
        </div>
      </div>
    </section>

    <!-- How It Works Section -->
    <section class="how-it-works">
      <h2>Как мы работаем</h2>
      <div class="steps">
        <div class="step">
          <div class="step-number">1</div>
          <h3>Оставьте заявку</h3>
          <p>Опишите ваш проект и пожелания</p>
        </div>
        <div class="step">
          <div class="step-number">2</div>
          <h3>Обсуждение</h3>
          <p>Наш менеджер свяжется с вами для уточнения деталей</p>
        </div>
        <div class="step">
          <div class="step-number">3</div>
          <h3>Реализация</h3>
          <p>Мы воплотим ваши идеи в жизнь</p>
        </div>
        <div class="step">
          <div class="step-number">4</div>
          <h3>Результат</h3>
          <p>Получите готовый проект и поддержку</p>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
// Можно добавить дополнительную логику если потребуется
</script>

<style scoped>
.general-view {
  width: 100%;
}

.hero {
  background: linear-gradient(135deg, #6366f1 0%, #a855f7 100%);
  color: white;
  padding: 6rem 2rem;
  text-align: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero h1 {
  font-size: 3.5rem;
  margin-bottom: 1rem;
  font-weight: bold;
}

.subtitle {
  font-size: 1.5rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

.cta-button {
  padding: 1rem 2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.3s ease;
}

.primary {
  background-color: white;
  color: #6366f1;
}

.primary:hover {
  background-color: #f3f4f6;
  transform: translateY(-2px);
}

.secondary {
  border: 2px solid white;
  color: white;
}

.secondary:hover {
  background-color: rgba(255, 255, 255, 0.1);
  transform: translateY(-2px);
}

.services, .how-it-works {
  padding: 4rem 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 3rem;
  color: #1f2937;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.service-card {
  padding: 2rem;
  text-align: center;
  background-color: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.service-card:hover {
  transform: translateY(-5px);
}

.service-card i {
  font-size: 2.5rem;
  color: #6366f1;
  margin-bottom: 1rem;
}

.service-card h3 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: #1f2937;
}

.service-card p {
  color: #6b7280;
  line-height: 1.6;
}

.steps {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
}

.step {
  text-align: center;
  padding: 2rem;
}

.step-number {
  width: 3rem;
  height: 3rem;
  background-color: #6366f1;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
  font-weight: bold;
  margin: 0 auto 1rem;
}

.step h3 {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
  color: #1f2937;
}

.step p {
  color: #6b7280;
}

@media (max-width: 768px) {
  .hero h1 {
    font-size: 2.5rem;
  }

  .subtitle {
    font-size: 1.25rem;
  }

  .cta-buttons {
    flex-direction: column;
  }

  .services-grid {
    grid-template-columns: 1fr;
  }
}
</style>