import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { Bookmark } from "lucide-react";
import { ReactElement } from "react";

export default function ItemFavouriteButton({ item }: { item: ItemExtended }): ReactElement {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button
            asChild
            className="size-8 p-0"
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
