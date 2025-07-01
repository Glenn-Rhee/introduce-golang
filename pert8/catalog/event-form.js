import { LitElement, html, css } from '/lit-core.min.js';

class EventForm extends LitElement {
    static properties = {
        event: { type: Object },
        isEdit: { type: Boolean }
    };

    constructor() {
        super();
        this.event = { 
            title: '', 
            organizer: '', 
            description: '', 
            date: '', 
            location: '', 
            price: 0,
            capacity: 0,
            image_url: '',
            status: 'upcoming'
        };
        this.isEdit = false;
    }

    static styles = css`
        :host {
            display: block;
            font-family: 'Roboto', 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #f9fafb;
            color: #111827;
            padding: 20px;
            box-sizing: border-box;
            min-height: 100vh;
        }
        h2 {
            text-align: center;
            font-size: 2rem;
            margin-bottom: 30px;
        }
        .form-container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            padding: 40px;
            border-radius: 12px;
            box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
        }
        .form-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            font-size: 1rem;
            margin-bottom: 8px;
            color: #374151;
        }
        input, textarea, select {
            width: 100%;
            padding: 12px 16px;
            border-radius: 8px;
            border: 1px solid #d1d5db;
            font-size: 1rem;
            box-sizing: border-box;
            transition: border-color 0.3s ease;
        }
        textarea {
            min-height: 120px;
            resize: vertical;
        }
        input:focus, textarea:focus, select:focus {
            outline: none;
            border-color: #3b82f6;
            box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
        }
        button {
            width: 100%;
            padding: 14px 20px;
            border-radius: 8px;
            background-color: #3b82f6;
            color: white;
            border: none;
            cursor: pointer;
            font-size: 1rem;
            font-weight: 500;
            transition: background-color 0.3s ease;
            margin-top: 10px;
        }
        button:hover {
            background-color: #2563eb;
        }
        .button-group {
            display: flex;
            gap: 10px;
        }
        .button-cancel {
            background-color: #9ca3af;
        }
        .button-cancel:hover {
            background-color: #6b7280;
        }
        @media (max-width: 640px) {
            .form-container {
                padding: 30px 20px;
            }
            button {
                font-size: 0.95rem;
            }
        }
    `;

    // Pre-populate form fields if editing an event
    async firstUpdated() {
        const eventId = this.getIdFromUrl();
        if (eventId) {
            this.isEdit = true;
            const response = await fetch(`http://localhost:28081/api/events/${eventId}`);
            if (response.ok) {
                const event = await response.json();
                // Format date for datetime-local input
                if (event.date) {
                    const date = new Date(event.date);
                    event.formattedDate = date.toISOString().slice(0, 16);
                }
                this.event = event;
            } else {
                console.error('Failed to fetch event');
            }
        }
    }

    getIdFromUrl() {
        const path = window.location.pathname;
        const segments = path.split('/');
        return segments[2];
    }

    async handleSubmit(e) {
        e.preventDefault();
        
        // Prepare event data
        const eventData = {...this.event};
        
        // Convert formatted date back to ISO string
        if (eventData.formattedDate) {
            eventData.date = new Date(eventData.formattedDate).toISOString();
            delete eventData.formattedDate;
        }
        
        const method = this.isEdit ? 'PUT' : 'POST';
        const endpoint = this.isEdit
            ? `http://localhost:28081/api/events/${this.event.id_event}`
            : 'http://localhost:28081/api/events';

        try {
            const response = await fetch(endpoint, {
                method,
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(eventData),
            });

            if (response.ok) {
                window.location.href = '/';
            } else {
                const errorData = await response.json();
                console.error('Error submitting form:', errorData);
                alert('Error saving event: ' + (errorData.message || 'Unknown error'));
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Network error');
        }
    }

    handleInput(e) {
        const { name, value, type } = e.target;
        
        // Special handling for different input types
        if (name === 'date') {
            this.event = { 
                ...this.event, 
                formattedDate: value 
            };
        } else {
            this.event = { 
                ...this.event, 
                [name]: type === 'number' ? parseFloat(value) : value 
            };
        }
    }

    render() {
        return html`
            <h2>${this.isEdit ? 'Edit Event' : 'Add New Event'}</h2>
            <div class="form-container">
                <form class="form" @submit="${this.handleSubmit}">
                    <div class="form-group">
                        <label for="title">Judul Event</label>
                        <input
                            type="text"
                            name="title"
                            id="title"
                            placeholder="Masukkan judul event"
                            .value="${this.event.title}"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="organizer">Penyelenggara</label>
                        <input
                            type="text"
                            name="organizer"
                            id="organizer"
                            placeholder="Masukkan nama penyelenggara"
                            .value="${this.event.organizer}"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="description">Deskripsi</label>
                        <textarea
                            name="description"
                            id="description"
                            placeholder="Masukkan deskripsi event"
                            .value="${this.event.description}"
                            @input="${this.handleInput}"
                        ></textarea>
                    </div>
                    <div class="form-group">
                        <label for="date">Tanggal dan Waktu</label>
                        <input
                            type="datetime-local"
                            name="date"
                            id="date"
                            .value="${this.event.formattedDate || ''}"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="location">Lokasi</label>
                        <input
                            type="text"
                            name="location"
                            id="location"
                            placeholder="Masukkan lokasi event"
                            .value="${this.event.location}"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="price">Harga Tiket</label>
                        <input
                            type="number"
                            name="price"
                            id="price"
                            placeholder="Masukkan harga tiket"
                            .value="${this.event.price}"
                            step="0.01"
                            min="0"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="capacity">Kapasitas</label>
                        <input
                            type="number"
                            name="capacity"
                            id="capacity"
                            placeholder="Masukkan kapasitas maksimum"
                            .value="${this.event.capacity}"
                            min="1"
                            @input="${this.handleInput}"
                            required
                        />
                    </div>
                    <div class="form-group">
                        <label for="image_url">URL Gambar</label>
                        <input
                            type="text"
                            name="image_url"
                            id="image_url"
                            placeholder="Masukkan URL gambar event"
                            .value="${this.event.image_url || ''}"
                            @input="${this.handleInput}"
                        />
                    </div>
                    <div class="form-group">
                        <label for="status">Status</label>
                        <select
                            name="status"
                            id="status"
                            .value="${this.event.status}"
                            @change="${this.handleInput}"
                            required
                        >
                            <option value="upcoming">Upcoming</option>
                            <option value="ongoing">Ongoing</option>
                            <option value="completed">Completed</option>
                            <option value="cancelled">Cancelled</option>
                        </select>
                    </div>
                    <div class="button-group">
                        <button type="button" class="button-cancel" @click="${() => window.location.href = '/'}">
                            Cancel
                        </button>
                        <button type="submit">
                            ${this.isEdit ? 'Update Event' : 'Add Event'}
                        </button>
                    </div>
                </form>
            </div>
        `;
    }
}

customElements.define('event-form', EventForm);