import React, { ReactElement, createContext, useState } from 'react';

export type TBreadCrumbsContext = {
  path: React.ReactElement[];
  setPath: React.Dispatch<React.SetStateAction<React.ReactElement[]>>;
};

const initialBreadCrumbs: TBreadCrumbsContext = {
  path: [],
  setPath: () => {},
};

export const BreadCrumbsContext = createContext<TBreadCrumbsContext>(initialBreadCrumbs);

export function BreadCrumbsContextProvider(props: {
  children: React.ReactElement | React.ReactElement[];
}): ReactElement {
  const [path, setPath] = useState<Array<React.ReactElement>>([<>Home</>, <>User</>]);

  return (
    <BreadCrumbsContext.Provider value={{ path, setPath }}>
      <>
        <div>
          <span>You are here: </span>
          {path.map((item) => (
            <span> / {item}</span>
          ))}
        </div>
        {props.children}
      </>
    </BreadCrumbsContext.Provider>
  );
}
