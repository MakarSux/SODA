<template>
  <div class="about-us">
    <!-- Company Info Section -->
    <section class="company-info">
      <div class="content-wrapper">
        <h1>О компании SODA Design</h1>
        <p class="intro">
          Мы - команда профессиональных дизайнеров, объединенных страстью к созданию
          уникальных визуальных решений. С 2020 года мы помогаем бизнесам и частным
          клиентам воплощать их идеи в жизнь.
        </p>
      </div>
    </section>

    <!-- Our Values Section -->
    <section class="values">
      <div class="content-wrapper">
        <h2>Наши ценности</h2>
        <div class="values-grid">
          <div class="value-card">
            <i class="fas fa-heart"></i>
            <h3>Качество</h3>
            <p>Мы не идем на компромиссы когда дело касается качества нашей работы</p>
          </div>
          <div class="value-card">
            <i class="fas fa-clock"></i>
            <h3>Пунктуальность</h3>
            <p>Ваше время ценно для нас, поэтому мы всегда соблюдаем сроки</p>
          </div>
          <div class="value-card">
            <i class="fas fa-lightbulb"></i>
            <h3>Креативность</h3>
            <p>Каждый проект уникален и заслуживает индивидуального подхода</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Team Section -->
    <section class="team">
      <div class="content-wrapper">
        <h2>Наша команда</h2>
        <div class="team-grid">
          <div class="team-member">
            <img src="https://via.placeholder.com/150" alt="Team Member">
            <h3>Анна Петрова</h3>
            <p class="position">Арт-директор</p>
            <p class="description">10+ лет опыта в графическом дизайне</p>
          </div>
          <div class="team-member">
            <img src="https://via.placeholder.com/150" alt="Team Member">
            <h3>Михаил Иванов</h3>
            <p class="position">Ведущий дизайнер</p>
            <p class="description">Специалист по UI/UX дизайну</p>
          </div>
          <div class="team-member">
            <img src="https://via.placeholder.com/150" alt="Team Member">
            <h3>Елена Сидорова</h3>
            <p class="position">Моушн-дизайнер</p>
            <p class="description">Эксперт по анимации и видео</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Contact Section -->
    <section class="contact">
      <div class="content-wrapper">
        <h2>Свяжитесь с нами</h2>
        <div class="contact-info">
          <div class="contact-item">
            <i class="fas fa-map-marker-alt"></i>
            <h3>Адрес</h3>
            <p>г. Москва, ул. Примерная, д. 123</p>
          </div>
          <div class="contact-item">
            <i class="fas fa-phone"></i>
            <h3>Телефон</h3>
            <p>+7 (999) 123-45-67</p>
          </div>
          <div class="contact-item">
            <i class="fas fa-envelope"></i>
            <h3>Email</h3>
            <p>info@sodadesign.ru</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
// Можно добавить дополнительную логику если потребуется
</script>

<style scoped>
.about-us {
  width: 100%;
}

.content-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.company-info {
  background: linear-gradient(135deg, #6366f1 0%, #a855f7 100%);
  color: white;
  padding: 6rem 0;
  text-align: center;
}

.company-info h1 {
  font-size: 3rem;
  margin-bottom: 2rem;
}

.intro {
  font-size: 1.25rem;
  max-width: 800px;
  margin: 0 auto;
  line-height: 1.8;
  opacity: 0.9;
}

section {
  padding: 4rem 0;
}

h2 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 3rem;
  color: #1f2937;
}

.values-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.value-card {
  text-align: center;
  padding: 2rem;
  background-color: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.value-card:hover {
  transform: translateY(-5px);
}

.value-card i {
  font-size: 2.5rem;
  color: #6366f1;
  margin-bottom: 1rem;
}

.value-card h3 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: #1f2937;
}

.value-card p {
  color: #6b7280;
  line-height: 1.6;
}

.team {
  background-color: #f3f4f6;
}

.team-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.team-member {
  text-align: center;
  background-color: white;
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.team-member img {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  margin-bottom: 1rem;
  object-fit: cover;
}

.team-member h3 {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
  color: #1f2937;
}

.position {
  color: #6366f1;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.description {
  color: #6b7280;
}

.contact-info {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.contact-item {
  text-align: center;
  padding: 2rem;
  background-color: white;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.contact-item i {
  font-size: 2rem;
  color: #6366f1;
  margin-bottom: 1rem;
}

.contact-item h3 {
  font-size: 1.25rem;
  margin-bottom: 0.5rem;
  color: #1f2937;
}

.contact-item p {
  color: #6b7280;
}

@media (max-width: 768px) {
  .company-info h1 {
    font-size: 2rem;
  }

  .intro {
    font-size: 1rem;
  }

  .values-grid,
  .team-grid,
  .contact-info {
    grid-template-columns: 1fr;
  }
}
</style>