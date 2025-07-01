-- Membuat database event_realm
CREATE DATABASE IF NOT EXISTS event_realm;

-- Menggunakan database event_realm
USE event_realm;

-- Membuat tabel events
CREATE TABLE IF NOT EXISTS events (
    id_event INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    organizer VARCHAR(255) NOT NULL,
    description TEXT,
    date DATETIME NOT NULL,
    location VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    capacity INT NOT NULL,
    image_url VARCHAR(255),
    status ENUM('upcoming', 'ongoing', 'completed', 'cancelled') DEFAULT 'upcoming',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Menambahkan data contoh ke tabel events
INSERT INTO events (title, organizer, description, date, location, price, capacity, image_url, status) VALUES
('BlizzCon 2025', 'Blizzard Entertainment', 'The ultimate gaming experience for Blizzard fans. Get the latest updates on World of Warcraft, Diablo IV, Overwatch 2, and more. Meet your favorite developers and join exclusive panels.', '2025-10-25 09:00:00', 'Anaheim Convention Center, California, USA', 3880000.00, 35000, 'https://s.yimg.com/ny/api/res/1.2/4cvYFudcaHqri9IA8w0TiA--/YXBwaWQ9aGlnaGxhbmRlcjt3PTk2MDtoPTU0MA--/https://media.zenfs.com/en/pc_gamer_708/9fce85f2e6bffb16a398554669210ba6', 'upcoming'),
('Tokyo Game Show 2025', 'Computer Entertainment Supplier''s Association', 'The world''s biggest gaming expo in Asia. Experience unreleased games, attend developer sessions, and witness world premieres of next-gen titles.', '2025-09-18 10:00:00', 'Makuhari Messe, Chiba, Japan', 2030000.00, 250000, 'https://events.nikkeibp.co.jp/tgs/2025/jp/exhibitor/ogp.png', 'upcoming'),
('Utaite Dream Festival', 'VocaSphere Productions', 'A two-day festival featuring top Japanese utaite and virtual singers. Join exclusive live performances from Mafumafu, Eve, Soraru, and special guest virtual performers.', '2025-05-15 17:30:00', 'Tokyo Dome, Tokyo, Japan', 1850000.00, 55000, 'https://img.youtube.com/vi/QMGvSE9CyaY/maxresdefault.jpg', 'upcoming'),
('World of Warcraft: Shadows Rising - Global Raid Tournament', 'ESL Gaming', 'The most prestigious WoW raid competition. Watch top guilds from around the world compete to clear the newest raid on Mythic difficulty for a prize pool of $500,000.', '2025-06-12 15:00:00', 'Lanxess Arena, Cologne, Germany', 1450000.00, 18000, 'placeholder.jpg', 'upcoming'),
('Dungeons & Dragons Epic: Crystal Caverns', 'Wizards of the Coast', 'An immersive D&D experience where hundreds of tables participate in the same adventure simultaneously. Join as a player or watch DM masters craft unforgettable stories.', '2025-07-20 10:00:00', 'Gen Con, Indianapolis, USA', 1015000.00, 2000, 'placeholder.jpg', 'upcoming'),
('Utaite World Connection', 'VTuber Entertainment Group', 'A digital concert featuring top utaite. Experience stunning AR performances and interactive elements never seen before.', '2025-08-05 19:00:00', 'Virtual Event - Global', 680000.00, 100000, 'placeholder.jpg', 'upcoming'),
('The International 2025: Dota 2 Championships', 'Valve Corporation', 'The world''s biggest esports tournament returns with the top Dota 2 teams competing for the Aegis of Champions and a record-breaking prize pool expected to exceed $40 million.', '2025-08-20 10:00:00', 'Mercedes-Benz Arena, Shanghai, China', 2000000.00, 18000, 'placeholder.jpg', 'upcoming'),
('Final Fantasy Music Concert: Distant Worlds', 'AWR Music Productions', 'Experience the beloved music of Final Fantasy performed by a full symphony orchestra and choir. Features compositions from the entire FF series, conducted by Arnie Roth with special guest Nobuo Uematsu.', '2025-04-22 20:00:00', 'Royal Albert Hall, London, UK', 1600000.00, 5500, 'placeholder.jpg', 'upcoming');

-- Menampilkan data dari tabel events
SELECT * FROM events;