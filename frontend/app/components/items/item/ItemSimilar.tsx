import { Item } from "@/proto/app/protos/items";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { marketplaces } from "@/store/categoryStore";
import { Marketplace } from "@/types/categoryTypes";
import { MessageCircleIcon, Star } from "lucide-react";
import { ReactElement } from "react";

export default function ItemSimilar({
  item,
  market,
}: {
  item: ItemExtended;
  market: number;
}): ReactElement {
  const marketUrl: string | undefined = marketplaces.find(
    (m: Marketplace): boolean => m.id === market,
  )?.value;

  if (!marketUrl) {
    return <></>;
  }

  return (
    <div className="text-sm">
      {item.similar.length > 0 ? <p className="font-bold">Похожие:</p> : null}
      {item.similar.map(
        (simItem: Item): ReactElement => (
          <a
            href={marketUrl + simItem.url}
            target="_blank"
            rel="noopener noreferrer"
            className="bold inline-flex flex-wrap items-center pb-2 hover:cursor-pointer hover:underline"
          >
            <span className="pe-1">{simItem.name}</span>

            {"["}

            {simItem.rating == 0 ? (
              <div className="inline-flex flex-wrap items-center">
                <Star className="size-3 fill-yellow-500 text-yellow-500" />
                <span>Нет отзывов</span>
              </div>
            ) : (
              <>
                <Star className="size-3 fill-yellow-500 text-yellow-500" />
                <span>{simItem.rating?.toFixed(1)}</span>
                <div className="inline-flex flex-wrap items-center">
                  <MessageCircleIcon className="size-3 fill-sky-300 text-sky-300" />
                  <span>{simItem.comments}</span>
                </div>
              </>
            )}
            {"]"}
          </a>
        ),
      )}
    </div>
  );
}
