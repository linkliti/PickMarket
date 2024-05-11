import { saveToFav } from "@/components/favorites/itemSaveLoad";
import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { useToast } from "@/components/ui/use-toast";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { Bookmark, TriangleAlert } from "lucide-react";
import { ReactElement } from "react";

export default function ItemFavouriteButton({ item }: { item: ItemExtended }): ReactElement {
  const [formPrefs] = useFilterStore((state: FilterStore) => [state.formPrefs]);

  const { toast } = useToast();

  function onClick(): void {
    if (!formPrefs) {
      toast({
        className: "p-4 border border-border",
        action: (
          <div className="flex w-full items-center">
            <TriangleAlert className="mr-2" />
            <span className="first-letter:capitalize">Не удалось сохранить товар</span>
          </div>
        ),
      });
      return;
    }
    saveToFav({ item, filters: formPrefs });
    toast({
      className: "p-4 border border-border",
      action: (
        <div className="flex w-full items-center">
          <TriangleAlert className="mr-2" />
          <span className="first-letter:capitalize">Товар добавлен в избранное</span>
        </div>
      ),
    });
  }

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button
            asChild
            className="size-8 p-0"
            onClick={onClick}
          >
            <div className="p-1">
              <Bookmark className="size-fit" />
            </div>
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>Сохранить товар и предпочтения в Избранное</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
