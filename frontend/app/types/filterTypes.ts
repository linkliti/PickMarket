import { UserPref } from "@/proto/app/protos/reqHandlerTypes";
import { StringList } from "@/proto/app/protos/types";

export type FiltersStore = {
  market: number,
  pageUrl: string
  numOfPages: number
  params: string
  userQuery: string
  prefs: Record<string, UserPref>
  setPageData: (market: number, pageUrl: string) => void;
  setNumOfPages: (numOfPages: number) => void;
  setParams: (params: string) => void;
  setUserQuery: (userQuery: string) => void;
  modifyPref: (key: string, value: UserPref) => void;
  resetStore: () => void;
};

export type FilterPrefValue =
  | {
      oneofKind: "numVal";
      numVal: number;
    }
  | {
      oneofKind: "listVal";
      listVal: StringList;
    }
  | {
      oneofKind: undefined;
    };