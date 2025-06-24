import { html, css, LitElement } from './lib/lit-core.min.js';

class NotesApp extends LitElement {
  static styles = css`
    :host {
      display: block;
      font-family: Arial, sans-serif;
      max-width: 500px;
      margin: auto;
    }
    textarea, input, button {
      width: 100%;
      margin: 5px 0;
      padding: 8px;
    }
    .note {
      border: 1px solid #ccc;
      padding: 10px;
      margin-top: 5px;
      background: #f9f9f9;
    }
    .delete-btn {
      background: red;
      color: white;
      border: none;
      padding: 5px;
      cursor: pointer;
    }
  `;

  static properties = {
    notes: { type: Array }
  };

  constructor() {
    super();
    this.notes = [];
  }

  connectedCallback() {
    super.connectedCallback();
    this.fetchNotes();
  }

  async fetchNotes() {
    const response = await fetch('/notes');
    this.notes = await response.json();
  }

  async addNote() {
    const title = this.shadowRoot.getElementById('title').value;
    const content = this.shadowRoot.getElementById('content').value;
    
    if (!title || !content) return;

    await fetch('/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, content })
    });

    this.fetchNotes();
  }

  async deleteNote(id) {
    if (!confirm("Are you sure you want to delete this note?")) {
      return;
    }
    
    await fetch(`/delete?id=${id}`, { method: 'DELETE' });
    this.fetchNotes();
  }

  render() {
    return html`
      <h2>Notes App</h2>
      <input id="title" type="text" placeholder="Title">
      <textarea id="content" placeholder="Write your note..."></textarea>
      <button @click="${this.addNote}">Add Note</button>
      
      ${this.notes.map(note => html`
        <div class="note">
          <h3>${note.title}</h3>
          <p>${note.content}</p>
          <button class="delete-btn" @click="${() => this.deleteNote(note.id)}">Delete</button>
        </div>
      `)}
    `;
  }
}

customElements.define('notes-app', NotesApp);
