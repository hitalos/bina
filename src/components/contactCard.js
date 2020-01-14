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
        otherTelephone: 'Outro'
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
          new RegExp(`(^|\\b)${el.className.split(' ').join('|')}(\\b|$)`, 'gi'),
          ' '
        )
      }
    },
  },
})
