import { ItemContext } from "@/components/items/ItemContext";
import ItemSimilar from "@/components/items/item/ItemSimilar";
import ItemTopPrefs from "@/components/items/item/ItemTopPrefs";
import { cn } from "@/lib/utils";
import { marketplaces } from "@/store/categoryStore";
import { MessageCircleIcon, Star } from "lucide-react";
import { ReactElement, useContext, useState } from "react";

export default function ItemDescription({ className = "" }: { className?: string }): ReactElement {
  const { item, market } = useContext(ItemContext);

  const rating: number = item.rating || 0;
  const comments: number = item.comments || 0;

  const [blurred, setBlurred] = useState(item.isAdult);

  return (
    <div
      className={cn(
        "flex items-start gap-4 max-sm:flex-col max-sm:items-center md:w-8/12",
        className,
      )}
    >
      <div className="flex flex-col items-center">
        <button
          type="button"
          className={cn(
            "w-[100px] cursor-default rounded-2xl bg-zinc-300",
            blurred && "cursor-pointer blur-sm hover:opacity-80",
          )}
          onClick={(): void => setBlurred(false)}
        >
          <img
            src={item.imageUrl}
            alt={item.name}
            className="h-full w-full rounded-2xl object-cover"
          />
        </button>
        {item.original && <span className="text-sm font-bold">Оригинал</span>}
      </div>
      <div className="flex w-full flex-col">
        <a
          href={marketplaces[market].value + item.url}
          target="_blank"
          rel="noopener noreferrer"
          className="pb-2 text-lg font-bold hover:cursor-pointer hover:underline"
        >
          {item.name}
        </a>
        <div className="inline-flex items-center gap-1 pb-2 font-bold">
          <Star className="size-4 fill-yellow-500 text-yellow-500" />
          {rating == 0 ? (
            <span>Нет отзывов</span>
          ) : (
            <>
              <span className="pr-2">{rating.toFixed(1)}</span>
              <MessageCircleIcon className="size-4 fill-sky-300 text-sky-300" />
              <span>{comments}</span>
            </>
          )}
        </div>
        <ItemSimilar />
        <ItemTopPrefs />
      </div>
    </div>
  );
}
