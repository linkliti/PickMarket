import WhiteBlock from "@/components/base/WhiteBlock";
import ItemDescription from "@/components/items/item/ItemDescription";
import ItemFavouriteButton from "@/components/items/item/ItemFavouriteButton";
import ItemWeightSection from "@/components/items/item/ItemWeightSection";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { ReactElement } from "react";

export default function ItemBlock({
  item,
  maxTotalWeight,
}: {
  item: ItemExtended;
  maxTotalWeight: number;
}): ReactElement {
  return (
    <WhiteBlock className="w-full">
      <div className="flex w-full justify-between gap-5 rounded-2xl max-md:max-w-full max-md:flex-wrap">
        <ItemDescription
          item={item}
          className="max-m:w-ful w-8/12"
        />
        <div className="flex w-4/12 flex-col max-md:w-full">
          <div className="inline-flex items-center text-2xl font-bold">
            <div className="inline-flex grow items-center justify-center gap-1">
              <h2>{item.item?.price ? `${item.item?.price} ₽` : "Нет информации о цене"}</h2>
              {item.item?.oldPrice ? (
                <span className="text-sm text-zinc-500"> ({item.item?.oldPrice} ₽)</span>
              ) : (
                ""
              )}
            </div>
            <ItemFavouriteButton item={item} />
          </div>
          <ItemWeightSection
            item={item}
            maxTotalWeight={maxTotalWeight}
            className=""
          />
        </div>
      </div>
    </WhiteBlock>
  );
}
