import { createApp, ref, onMounted, onUnmounted } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.js'

createApp({
	setup() {
		const displaySection = ref('software')
		
		function updateSection(){
			const path = window.location.pathname;
			const section = path.split('/').pop();
			displaySection.value = section || 'software'
		}

		updateSection()
		return {
			displaySection, updateSection
		}
	}
}).mount("#app")
