import { json, select } from 'd3'
import i18n from './i18n.mjs'

const results = select('#results')

const normalize = (str) => str.toLowerCase().normalize('NFD').replace(/[\u0300-\u036f]/g, '')

const filter = (param) => (d) => {
	if (d.others?.title && normalize(d.others.title).includes(param)) return true
	if (d.others?.department && normalize(d.others.department).includes(param)) return true
	if (d.others?.physicalDeliveryOfficeName && normalize(d.others.physicalDeliveryOfficeName).includes(param)) return true
	if (d.phones && Object.keys(d.phones).some((k) => normalize(d.phones[k]).includes(param))) return true
	if (d.emails && Object.keys(d.emails).some((k) => normalize(d.emails[k]).includes(param))) return true
	if (normalize(d.fullName).includes(param)) return true
	if (normalize(d.fullName).includes(param)) return true
	if (d.others?.employeeID && normalize(d.others.employeeID).includes(param)) return true

	return false
}

const cardTemplate = ({ id, fullName, emails, others, phones }) => {
	const dp = others?.department || ''
	const pdon = others?.physicalDeliveryOfficeName || ''
	const title = others?.title || ''
	const email = emails?.mail || emails?.proxyAddresses || ''
	const eID = others?.employeeID || ''
	const mainPhone = phones.ipPhone || phones.mobile || phones.telephoneNumber || phones.homePhone || phones.otherTelephone || phones.facsimileTelephoneNumber || ''

	return `
		<div class="front">
			<header>
				<h2>${fullName}</h2>
				<img src="${id ? `/contacts/photo/${id}` : '/images/logo'}"/>
			</header>
			<main>
				${dp || pdon ? `Lotação: <strong>${dp}${pdon && dp ? ' - ' : ''}${pdon}</strong>` : ''}
			</main>
			<footer>
				<a href="tel:${mainPhone}">${mainPhone}</a>
			</footer>
		</div>
		<div class="back">
			<header>${fullName}</header>
			<main>
				<dl>
					${title ? `<dt>Vínculo:</dt><dd>${title}</dd>` : ''}
					${email ? `<dt>Email:</dt><dd>${email}</dd>` : ''}
					${eID ? `<dt>Matrícula:</dt><dd>${eID}</dd>` : ''}
					${Object.keys(phones).map((k) => `<dt>${i18n['pt-BR'][k]}:</dt><dd>${phones[k]}</dd>`).join('')}
				</dl>
			</main>
			<footer>
				${id ? `<a href="/contacts/vcard/${id}">Baixar Vcard</a>` : ''}
				${id ? `<a class="modal" data-url="/contacts/qrcode/${id}">Ver QRcode</a>` : ''}
			</footer>
		</div>`
}

const showQRcode = (url) => {
	const modal = document.createElement('dialog')
	modal.innerHTML = `
		<div class="modal-content">
			<span class="close-button">×</span>
			<img src="${url}" alt="QR Code"/>
		</div>`
	document.body.appendChild(modal)
	const closeButton = modal.querySelector('.close-button')

	closeButton.addEventListener('click', () => {
		modal.close()
		document.body.removeChild(modal)
	})
	window.addEventListener('click', (ev) => {
		if (ev.target === modal) {
			modal.close()
			document.body.removeChild(modal)
		}
	})

	modal.showModal()
}

window.addEventListener('load', () => {
	json('/contacts/all.json').then((data) => {
		select('#search').on('keyup', (ev) => {
			const param = normalize(ev.currentTarget.value)
			if (param.length >= 3) {
				const filtered = data.filter(filter(param))
				select('body > main > p').text(`Encontrado(s) ${filtered.length} contato(s).`)
				const cards = results.selectAll('div.card')
					.data(filtered, (d) => d.id || d.fullName)

				cards.enter()
					.append('div')
					.classed('card', true)
					.classed('contact', (d) => d.objectClass === 'contact')
					.classed('user', (d) => d.objectClass === 'user')
					.html(cardTemplate)
					.on('click', (ev) => {
						if (ev.target.classList.contains('modal')) {
							ev.preventDefault()
							showQRcode(ev.target.dataset.url)
							return
						}
						ev.currentTarget.classList.toggle('turned')
					})

				cards.exit()
					.transition().duration(500)
					.style('opacity', 0)
					.remove()
			}
		})
	}).catch(console.error)
})
