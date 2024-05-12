import { ItemContext } from "@/components/items/ItemContext";
import { Item } from "@/proto/app/protos/items";
import { marketplaces } from "@/store/categoryStore";
import { MessageCircleIcon, Star } from "lucide-react";
import { ReactElement, useContext } from "react";

export default function ItemSimilar(): ReactElement {
  const { market, similar } = useContext(ItemContext);

  return (
    <div className="text-sm">
      {similar.length > 0 ? <p className="font-bold">Похожие:</p> : null}
      {similar.map(
        (simItem: Item): ReactElement => (
          <a
            href={marketplaces[market].value + simItem.url}
            key={simItem.url}
            target="_blank"
            rel="noopener noreferrer"
            className="bold inline-flex flex-wrap items-center pb-2 pr-2 hover:cursor-pointer hover:underline"
          >
            <span className="pe-1">
              {simItem.name} {simItem.original && " (Оригинал)"}
            </span>
            <div className="inline-flex items-center gap-1">
              {"["}
              {simItem.rating == 0 || !simItem.rating ? (
                <>
                  <Star className="size-3 fill-yellow-500 text-yellow-500" />
                  <span>Нет отзывов</span>
                </>
              ) : (
                <>
                  <Star className="size-3 fill-yellow-500 text-yellow-500" />
                  <span>{simItem.rating?.toFixed(1)}</span>
                  <MessageCircleIcon className="size-3 fill-sky-300 text-sky-300" />
                  <span>{simItem.comments}</span>
                </>
              )}
              {"]"}
            </div>
          </a>
        ),
      )}
    </div>
  );
}
