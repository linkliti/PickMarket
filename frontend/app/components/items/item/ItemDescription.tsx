import ItemTopPrefs from "@/components/items/item/ItemTopPrefs";
import { cn } from "@/lib/utils";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { MessageCircleIcon, Star } from "lucide-react";

export default function ItemDescription({
  item,
  className = "",
}: {
  item: ItemExtended;
  className: string;
}) {
  const rating: number = item.item?.rating || 0;
  const comments: number = item.item?.comments || 0;

  return (
    <div className={cn("flex items-start gap-4", className)}>
      <img
        src={item.item?.imageUrl}
        alt={item.item?.name}
        className="w-[100px] shrink-0 rounded-2xl border bg-zinc-300"
      />
      <div className="flex w-full flex-col">
        <a
          href={"https://ozon.ru" + item.item?.url}
          target="_blank"
          rel="noopener noreferrer"
          className="pb-2 text-lg font-bold hover:cursor-pointer hover:underline"
        >
          {item.item?.name}
        </a>
        <div className="inline-flex items-center gap-1 pb-2 font-bold">
          <Star className="size-4 fill-yellow-500 text-yellow-500" />
          <span className="pr-2">{rating === 0 ? "Нет рейтинга" : `${rating.toFixed(1)}`}</span>
          <MessageCircleIcon className="size-4 fill-sky-300 text-sky-300" />
          <span>{comments === 0 ? "Нет комментариев" : `${comments}`}</span>
        </div>
        <div className="text-sm">
          {item.similar.length > 0 ? <p className="font-bold">Похожие:</p> : null}
          {item.similar.map((item) => (
            <p>
              <a
                href={"https://ozon.ru" + item.url}
                target="_blank"
                rel="noopener noreferrer"
                className="bold pb-2 hover:cursor-pointer hover:underline"
              >
                {item.name}
              </a>
            </p>
          ))}
        </div>

        <ItemTopPrefs chars={item.chars} />
      </div>
    </div>
  );
}
