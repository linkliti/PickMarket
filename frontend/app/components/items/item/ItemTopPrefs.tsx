import { ItemContext } from "@/components/items/ItemContext";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { Characteristic } from "@/proto/app/protos/items";
import { MinusSquareIcon, PlusSquareIcon } from "lucide-react";
import { ReactElement, useContext, useState } from "react";

function printChar(char: Characteristic) {
  if (char.value.oneofKind === "listVal") {
    const listValStr = char.value.listVal.values.map((value) => value).join(", ");
    if (listValStr) {
      return listValStr;
    }
  } else if (char.value.oneofKind === "numVal") {
    return char.value.numVal.toString();
  }
  return "";
}

export default function ItemTopPrefs(): ReactElement {
  const { chars } = useContext(ItemContext);

  const sortedChars: Characteristic[] = [...chars].sort(
    (a: Characteristic, b: Characteristic): number => b.charWeight - a.charWeight,
  );

  const topChars: Characteristic[] = [];
  const weightedChars: Characteristic[] = [];
  const badChars: Characteristic[] = [];
  const nonWeightedChars: Characteristic[] = [];

  sortedChars.forEach((charItem: Characteristic): void => {
    const weightPercentage: number = (charItem.charWeight / charItem.maxWeight) * 100;
    if (weightPercentage >= 90) {
      topChars.push(charItem);
    } else if (weightPercentage < 10) {
      badChars.push(charItem);
    } else if (charItem.charWeight > 0) {
      weightedChars.push(charItem);
    } else {
      nonWeightedChars.push(charItem);
    }
  });

  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <table className="w-full table-auto">
        <thead></thead>
        <tbody>
          <TableChars
            chars={topChars}
            plusIcon={<PlusSquareIcon className="size-4 fill-green-400" />}
          />
          <TableChars
            chars={badChars}
            plusIcon={<MinusSquareIcon className="size-4 fill-red-400" />}
          />
        </tbody>
      </table>

      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
        className=""
      >
        <CollapsibleTrigger asChild>
          <p className="font-bold hover:cursor-pointer hover:underline">
            {isOpen ? "Скрыть" : "Показать все"}
          </p>
        </CollapsibleTrigger>
        <CollapsibleContent>
          <table className="w-full table-auto">
            <thead></thead>
            <tbody>
              <TableChars
                chars={weightedChars}
                plusIcon={<PlusSquareIcon className="size-4 fill-yellow-400" />}
              />
              <TableChars
                chars={nonWeightedChars}
                plusIcon={<MinusSquareIcon className="size-4 " />}
              />
            </tbody>
          </table>
        </CollapsibleContent>
      </Collapsible>
    </>
  );
}

function TableChars({
  chars,
  plusIcon: iconElement,
}: {
  chars: Characteristic[];
  plusIcon: ReactElement;
}): ReactElement {
  return (
    <>
      {chars.map(
        (char: Characteristic): ReactElement => (
          <tr
            key={char.key}
            className="flex flex-row gap-2 border-b border-b-gray-200"
          >
            <td className="flex w-3/5 items-baseline">
              {iconElement}
              <p className="w-1/2 grow text-wrap ps-2">{char.name + ": "}</p>
              {char.maxWeight > 0 && (
                <span className="my-auto text-xs font-bold text-zinc-500">
                  ({char.charWeight.toFixed(1)})
                </span>
              )}
            </td>
            <td className="flex w-2/5 items-center gap-2 text-wrap">
              <span className="r-0">{printChar(char)}</span>
            </td>
          </tr>
        ),
      )}
    </>
  );
}
