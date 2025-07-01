import { LitElement, html } from '/lit-core.min.js';
import '/event-catalog.js';
import '/event-form.js';

class AppRouter extends LitElement {
    static properties = {
        route: { type: String },
        eventId: { type: String }
    };

    constructor() {
        super();
        this.updateRoute();
        window.onpopstate = () => {
            this.updateRoute();
        };
    }

    updateRoute() {
        const path = window.location.pathname.replace(/^\/+/, '');
        const parts = path.split('/');
        this.route = parts[0] || 'catalog';
        this.eventId = parts[1] || null;
        this.requestUpdate();
    }

    navigateTo(route) {
        window.history.pushState({}, '', route);
        this.updateRoute();
    }

    render() {
        return html`
            ${this.route === 'catalog' || this.route === '' ? html`
                <event-catalog .navigateTo="${this.navigateTo.bind(this)}"></event-catalog>
            ` : this.route === 'add' ? html`
                <event-form></event-form>
            ` : this.route === 'edit' && this.eventId ? html`
                <event-form .eventId="${this.eventId}" .isEdit="${true}"></event-form>
            ` : html`
                <div style="text-align: center; padding: 50px; font-family: 'Roboto', sans-serif;">
                    <h2>404 - Halaman Tidak Ditemukan</h2>
                    <p>Maaf, halaman yang Anda cari tidak ditemukan.</p>
                    <button 
                        @click="${() => this.navigateTo('/')}"
                        style="background-color: #3b82f6; color: white; border: none; border-radius: 8px; padding: 12px 24px; cursor: pointer; margin-top: 20px;"
                    >
                        Kembali ke Halaman Utama
                    </button>
                </div>
            `}
        `;
    }
}

customElements.define('app-router', AppRouter);