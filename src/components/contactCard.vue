<template>
	<div class="flip-container" ref="flip-container">
		<div class="flipper">
			<div class="front">
				<md-layout :class="contact.objectClass">
					<md-card class="md-flex-100 md-with-hover">
						<md-card-header :title="contact.others && contact.others.title ? &quot;Vínculo: &quot; + contact.others.title : &quot;&quot;">
							<md-card-header-text>
								<div class="md-headline">{{ contact.fullName }}</div>
							</md-card-header-text>
							<md-card-media v-if="!contact.id">
								<img src="/images/logo.png">
							</md-card-media>
							<md-card-media v-else>
								<img v-if="contact.photo &amp;&amp; contact.photo.type" :src="'data:'+ contact.photo.type +';base64,' + contact.photo.data" alt="Foto">
								<img v-else :src="'/contacts/' + contact.id + '.jpg'" alt="Foto">
							</md-card-media>
						</md-card-header>
						<md-card-content>
							<p v-if="contact.others && contact.others.department">Lotação: <strong>{{ contact.others.department }} <span v-if="contact.others && contact.others.physicalDeliveryOfficeName">- {{ contact.others.physicalDeliveryOfficeName }}</span></strong></p>
						</md-card-content>
						<md-card-actions>
							<md-button :title="defaultPhone">
								<a class="md-display-1" :href="'tel:' + defaultPhone">{{ defaultPhone }}</a>
							</md-button>
							<div class="md-flex"></div>
							<md-button @click.native="invert">Ver mais detalhes</md-button>
						</md-card-actions>
					</md-card>
				</md-layout>
			</div>
			<div class="back">
				<md-layout :class="contact.objectClass">
					<md-card class="md-flex-100 md-with-hover">
						<md-card-content>
							<p v-if="contact.others && contact.others.title">Vínculo: <strong>{{ contact.others.title }}</strong></p>
							<div v-if="contact.emails">
								<p v-for="(email, key) in contact.emails" :title="key" :key="key">Mail: <strong>{{ email }}</strong></p>
							</div>
							<p v-if="contact.others && contact.others.employeeID">Matrícula: <strong>{{ contact.others.employeeID }}</strong></p>
							<p v-for="(phone, key) in contact.phones" :key="key">
								<span v-if="Array.isArray(phone)">
									<p v-for="num in phone" :key="num">{{ labels[key] }}: <strong>{{ num }}</strong></p>
								</span>
								<span v-else>{{ labels[key] }}: <strong>{{ phone }}</strong></span>
							</p>
						</md-card-content>
						<md-card-actions>
							<md-button v-if="contact.objectClass=='user'">
								<a :href="'/contacts/' + contact.id + '.vcf'">Baixar vCard</a>
							</md-button>
							<div class="md-flex"></div>
							<md-button @click.native="revert">voltar</md-button>
						</md-card-actions>
					</md-card>
				</md-layout>
			</div>
		</div>
	</div>
</template>

<script>
import Vue from 'vue'

export default Vue.component('contact-card', {
  props: ['contact'],
  data() {
    return {
      labels: {
        telephoneNumber: 'Principal',
        mobile: 'Celular',
        ipPhone: 'VoIP',
        facsimileTelephoneNumber: 'Fax',
        homePhone: 'Casa',
        otherTelephone: 'Outro',
      },
    }
  },
  computed: {
    defaultPhone() {
      const opts = [
        'ipPhone',
        'telephoneNumber',
        'facsimileTelephoneNumber',
        'mobile',
        'homePhone',
        'otherTelephone',
      ]
      const phones = Object.keys(this.contact.phones)
      const defaultPhone = opts.filter((label) => phones.indexOf(label) !== -1)[0]
      return this.contact.phones[defaultPhone]
    },
  },
  methods: {
    invert() {
      const el = this.$refs['flip-container']
      if (el.classList) el.classList.add('reverse')
      else el.className += ' reverse'
    },
    revert() {
      const el = this.$refs['flip-container']
      if (el.classList) el.classList.remove('reverse')
      else {
        el.className = el.className.replace(
          new RegExp(`(^|\\b)${el.className.split(' ').join('|')}(\\b|$)`, 'gi'),
          ' ',
        )
      }
    },
  },
})
</script>
