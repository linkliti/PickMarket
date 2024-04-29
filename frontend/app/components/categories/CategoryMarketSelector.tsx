import BodyHeader from "@/components/base/BodyHeader";
import { useState } from "react";

interface Marketplace {
  name: string;
  url: string;
}

const marketplaces: Marketplace[] = [
  { name: 'OZON', url: 'https://ozon.ru' },
  { name: 'Я.Маркет', url: 'https://market.yandex.ru' },
];

const categories = ['Элетроника', 'Одежда'];

export default function CategoryMarketSelector() {
  const [selectedMarketplace, setSelectedMarketplace] = useState<Marketplace | null>(null);
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null);
  const [text, setText] = useState<string>('');

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (selectedMarketplace && selectedCategory) {
      console.log({ selectedMarketplace, selectedCategory, text });
    }
  };

  return (
    <BodyHeader>
      <form onSubmit={handleSubmit}>
        <select value={selectedMarketplace?.name || ''} onChange={(e) => setSelectedMarketplace(marketplaces.find(m => m.name === e.target.value) || null)} required>
          <option value="">Выберите магазин</option>
          {marketplaces.map(m => (
            <option key={m.name} value={m.name}>{m.name}</option>
          ))}
        </select>
        <select
          value={selectedCategory || ''}
          onChange={(e) => setSelectedCategory(e.target.value || null)}
          required
        >
          <option value="">Выберите категорию</option>
          {categories.map(c => (
            <option key={c} value={c}>{c}</option>
          ))}
        </select>
        <input
          type="text"
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Поисковая фраза"
        />
        <button type="submit">Поиск</button>
      </form>

    </BodyHeader>
  );
}
