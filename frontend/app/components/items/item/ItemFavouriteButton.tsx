import { removeFromFav, saveToFav } from "@/components/favorites/itemSaveLoad";
import { FilterContext } from "@/components/items/FilterContext";
import { ItemContext } from "@/components/items/ItemContext";
import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { useToast } from "@/components/ui/use-toast";
import { cn } from "@/lib/utils";
import { Bookmark, CircleCheck } from "lucide-react";
import { ReactElement, useContext, useState } from "react";

export default function ItemFavouriteButton(): ReactElement {
  const { item, chars, similar, favUrl, isFav, market } = useContext(ItemContext);
  const { categoryURL, form } = useContext(FilterContext);
  const [isFavorite, setIsFavorite] = useState<boolean>(isFav);

  const { toast } = useToast();

  function addItemToFavs(): void {
    saveToFav(
      {
        item: {
          item: item,
          chars: chars,
          similar: similar,
          totalWeight: 0,
        },
        filters: form,
        categoryURL: categoryURL,
        market: market,
      },
      favUrl,
    );
    setIsFavorite(true);
    toast({
      className: "p-4 border border-border",
      action: (
        <div className="flex w-full items-center">
          <CircleCheck className="mr-2" />
          <span className="first-letter:capitalize">Товар добавлен в избранное</span>
        </div>
      ),
      duration: 3000,
    });
  }

  function removeItemFromFavs(): void {
    removeFromFav(favUrl);
    setIsFavorite(false);
    toast({
      className: "p-4 border border-border",
      action: (
        <div className="flex w-full items-center">
          <CircleCheck className="mr-2" />
          <span className="first-letter:capitalize">Товар убран из избранного</span>
        </div>
      ),
      duration: 3000,
    });
  }

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button
            asChild
            className="size-8 p-0"
            onClick={isFavorite ? removeItemFromFavs : addItemToFavs}
          >
            <div className="p-1">
              <Bookmark className={cn("size-fit", isFavorite && "fill-red-500")} />
            </div>
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>
            {isFavorite ? "Удалить из Избранного" : "Сохранить товар и предпочтения в Избранное"}
          </p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
