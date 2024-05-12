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

const charPercent = 0.9;

export default function ItemTopPrefs(): ReactElement {
  const { chars } = useContext(ItemContext);

  const weightedChars = chars.filter(
    (char) => char.charWeight >= char.maxWeight * charPercent && char.maxWeight > 0,
  );
  const nonWeightedChars = chars.filter(
    (char) => char.charWeight < char.maxWeight * charPercent || char.maxWeight === 0,
  );

  const [isOpen, setIsOpen] = useState(false);
  return (
    <>
      <table className="w-full table-auto">
        <thead></thead>
        <tbody>
          <TableChars
            chars={weightedChars}
            plusIcon={<PlusSquareIcon className="size-4 fill-green-400" />}
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
