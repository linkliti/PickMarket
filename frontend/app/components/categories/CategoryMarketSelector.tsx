import { useState } from "react";
// @ts-expect-error missing types
import { terminal } from "virtual:terminal";
import BodyHeader from "../base/BodyHeader";

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
              className={`justify-center rounded-2xl px-9 py-3.5 max-md:px-5 ${selectedMarketplace === id ? "pm-btn pm-btn-active" : "pm-btn"}`}
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
          <input
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
        <button
          className="justify-center rounded-2xl bg-sky-200 px-12 py-3.5 text-center max-md:px-5"
          onClick={showProducts}
        >
          Показать товары
        </button>
      </div>
    </BodyHeader>
  );
}
