<template>
  <div class="services-grid">
    <div class="service-banner" v-for="(banner, index) in banners" :key="index" :id="(index + 1)">
      
      <div class="service_image" v-if="(banner.image != '')">
        <img :src="banner.image" :alt="banner.title">
      </div>
      
      <div class="service-text">
        <h3>{{ banner.title }}</h3>
        <p>{{ banner.description }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import MaskOne from '@/assets/1m.png'
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
          image: MaskOne,
          title: 'Сайты под ключ',
          description: 'Создание полноценного сайта с нуля'
        },
        {
          image: '',
          title: 'Мобильные приложения',
          description: 'Разработка приложений для iOS и Android'
        },
        {
          image: '',
          title: 'UI / UX Дизайн',
          description: 'Создание привлекательных интерфейсов'
        },
        {
          image: '',
          title: 'SEO Оптимизация',
          description: 'Продвижение вашего сайта'
        },
        {
          image: '',
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
  grid-template-columns: 1fr;
  gap: 2rem;
  padding: 2rem 0;
}

.service-banner {
  position: relative;
  background: rgb(255, 0, 72);
  background: linear-gradient(90deg, rgba(43, 43, 43, 1) 35%, rgba(255, 0, 72, 1) 100%);
  border-radius: 10px;
  padding: 1.5rem;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.service-banner:hover {
  transform: translateY(-5px);
}


.service-banner .service-text {
  position: absolute;
  top: 0;
  left: 0;
}

.service-banner .service-text h3 {
  color: #ffffff;
  font-weight: 600;
  font-size: 50px;
  padding: 1em 0 0 0.5em;
}

.service-banner .service-text p {
  color: #ffffff;
  font-weight: 600;
  font-size: 20px;
  padding-left: 1em;
}

.services-grid div[id='1'] {
  position: relative;
  width: 100%;
  padding-top: 20.9%;
  box-sizing: border-box;
}

.services-grid div[id='1'] .service_image {
  position: absolute;
  top: 0;
  left: 0;
}

.services-grid div[id='1'] .service_image img {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
}


div[id='2'],
div[id='3'],
div[id='4'],
div[id='5']{
  min-height: 250px;
}
</style>
