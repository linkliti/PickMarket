import WhiteBlock from "@/components/base/WhiteBlock";
import ItemBlockFav from "@/components/favorites/ItemBlockFav";
import { clearFav, getFav } from "@/components/favorites/favSaveLoad";
import { FavRecord, ItemWithFilters } from "@/components/favorites/types";
import { Button } from "@/components/ui/button";
import { ReactElement, useEffect, useState } from "react";

export default function Favorites(): ReactElement {
  const [favorites, setFavorites] = useState<FavRecord>({});

  useEffect((): void => {
    const favs: FavRecord = getFav();
    setFavorites(favs);
  }, []);

  if (Object.keys(favorites).length === 0) {
    return <WhiteBlock className="w-full">Нет избранных товаров</WhiteBlock>;
  }

  return (
    <>
      <WhiteBlock className="inline-flex w-full">
        <h2 className="text-2xl font-bold">
          Просмотр избранного ({Object.keys(favorites).length})
        </h2>
        <Button
          className="ml-auto"
          onClick={(): void => {
            clearFav();
            setFavorites({});
          }}
        >
          Очистить
        </Button>
      </WhiteBlock>
      {Object.entries(favorites).map(
        ([key, val]: [string, ItemWithFilters]): ReactElement => (
          <ItemBlockFav
            item={val.item}
            key={key}
            market={0}
          />
        ),
      )}
    </>
  );
}
