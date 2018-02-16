/* eslint no-unused-vars: 0 no-undef: 0 */
const ContactCard = Vue.component('contact-card', {
  props: ['contact'],
  data() {
    return {
      labels: {
        telephoneNumber: 'Principal',
        mobile: 'Celular',
        ipPhone: 'VoIP',
        facsimileTelephoneNumber: 'Fax',
        homePhone: 'Casa',
        otherTelephone: 'Outro'
      },
    }
  },
  computed: {
    defaultPhone() {
      return this.contact.phones[Object.keys(this.contact.phones)[0]]
    },
  },
  template: '#contact-card-template',
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
          new RegExp(`(^|\\b)${className.split(' ').join('|')}(\\b|$)`, 'gi'),
          ' '
        )
      }
    },
  },
})
