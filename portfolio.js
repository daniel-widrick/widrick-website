import { createApp, ref, onMounted, onUnmounted } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.js'

createApp({
	setup() {
		const displaySection = ref('Software')
		
		function updateSection(){
			displaySection.value = window.location.hash.replace("#","") || "Software"
		}

		window.addEventListener("hashchange", updateSection);
		updateSection()
		return {
			displaySection, updateSection
		}
	}
}).mount("#app")
