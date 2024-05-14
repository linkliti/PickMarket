import WhiteBlock from "@/components/base/WhiteBlock";
import { ItemContext } from "@/components/items/ItemContext";
import ItemDescription from "@/components/items/item/ItemDescription";
import ItemFavouriteButton from "@/components/items/item/ItemFavouriteButton";
import ItemGoToSavedFilters from "@/components/items/item/ItemGoToSavedFilters";
import { ReactElement, useContext } from "react";

export default function ItemBlockFav(): ReactElement {
  const { item } = useContext(ItemContext);

  return (
    <WhiteBlock className="w-full">
      <div className="flex w-full justify-between gap-5 rounded-2xl max-md:max-w-full max-md:flex-wrap">
        <ItemDescription />
        <div className="flex w-4/12 flex-col max-md:w-full">
          <div className="inline-flex items-center text-2xl font-bold">
            <div className="inline-flex grow items-center justify-center gap-1">
              <h2>{item.price ? `${item.price} ₽` : "Нет информации о цене"}</h2>
              {item.oldPrice ? (
                <span className="text-sm text-zinc-500"> ({item.oldPrice} ₽)</span>
              ) : (
                ""
              )}
            </div>
            <ItemFavouriteButton />
          </div>
          <ItemGoToSavedFilters />
        </div>
      </div>
    </WhiteBlock>
  );
}
