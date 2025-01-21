<template>
  <div class="services-grid">
    <div class="service-banner" v-for="(banner, index) in banners" :key="index">
      <img :src="banner.image" :alt="banner.title">
      <h3>{{ banner.title }}</h3>
      <p>{{ banner.description }}</p>
    </div>
  </div>
</template>

<script>
import maskGroup1 from '@/assets/Mask group-1.png'
import maskGroup from '@/assets/Mask group.png'
import rectangle81 from '@/assets/Rectangle 81.png'
import rectangle82 from '@/assets/Rectangle 82.png'
import botBanner from '@/assets/Создание бота для раличных платформ.png'

export default {
  name: 'BannerService',
  data() {
    return {
      banners: [
        {
          image: maskGroup,
          title: 'Веб-разработка',
          description: 'Создание современных веб-приложений'
        },
        {
          image: maskGroup1,
          title: 'Мобильные приложения',
          description: 'Разработка приложений для iOS и Android'
        },
        {
          image: rectangle81,
          title: 'UI/UX Дизайн',
          description: 'Создание привлекательных интерфейсов'
        },
        {
          image: rectangle82,
          title: 'SEO Оптимизация',
          description: 'Продвижение вашего сайта'
        },
        {
          image: botBanner,
          title: 'Разработка ботов',
          description: 'Создание ботов для различных платформ'
        }
      ]
    }
  }
}
</script>

<style scoped>
.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  padding: 2rem 0;
}

.service-banner {
  background: #fff;
  border-radius: 10px;
  padding: 1.5rem;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.service-banner:hover {
  transform: translateY(-5px);
}

.service-banner img {
  width: 100%;
  height: auto;
  border-radius: 8px;
  margin-bottom: 1rem;
}

.service-banner h3 {
  margin: 1rem 0;
  color: #333;
}

.service-banner p {
  color: #666;
  font-size: 0.9rem;
}
</style>
