import BodyHeader from "@/components/base/BodyHeader";
import { Input } from "@/components/ui/input";
import { useState } from "react";

// @ts-expect-error missing types
import { terminal } from "virtual:terminal";

const marketplaces = {
  OZON: ["OZON", "https://ozon.ru"],
  YAND: ["Я.Маркет", "https://market.yandex.ru"],
};

export default function CategoryMarketSelector() {
  const [selectedMarketplace, setSelectedMarketplace] = useState("OZON");
  const [searchText, setSearchText] = useState("");

  const showProducts = () => {
    terminal.log(`Search Query: ${searchText}`);
    terminal.log(`Selected Marketplace ID: ${selectedMarketplace}`);
  };

  return (
    <BodyHeader>
      <div className="flex w-full gap-5 whitespace-nowrap font-bold max-lg:flex-wrap max-md:max-w-full">
        <div className="my-auto flex-auto text-2xl text-black">Маркетплейс:</div>
        <div className="flex gap-4 text-center text-lg">
          {Object.entries(marketplaces).map(([id, [name]]) => (
            <label
              key={id}
              className={`pm-btn ${selectedMarketplace === id ? "pm-btn-act" : "pm-btn-deact"}`}
            >
              <input
                type="radio"
                name="marketplace"
                className="appearance-none"
                value={id}
                checked={selectedMarketplace === id}
                onChange={() => setSelectedMarketplace(id)}
              />
              {name}
            </label>
          ))}
        </div>
        <div className="flex justify-between gap-5">
          <div className="my-auto text-center text-lg text-black">Поиск:</div>
          <Input
            type="text"
            className="justify-center rounded-2xl border border-solid border-black p-4 text-base"
            placeholder="Поиск"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
          />
        </div>
      </div>
      <div className="mt-5 flex gap-5 text-lg font-bold text-black max-md:max-w-full max-md:flex-wrap">
        <div className="my-auto flex-auto max-md:max-w-full">
          Выбранная категория: “Категория - ДочерКатегория”
        </div>
        <button className="pm-btn pm-btn-act" onClick={showProducts}>
          Показать товары
        </button>
      </div>
    </BodyHeader>
  );
}
