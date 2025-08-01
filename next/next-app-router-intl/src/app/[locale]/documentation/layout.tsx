import React, { ReactNode } from "react";
import Toc from "@/components/documentation/Toc/Toc";
import styles from "./documentation.module.css";

type DocumentationLayoutProps = {
  children: ReactNode;
};

const DocumentationLayout = ({ children }: DocumentationLayoutProps) => {


  return (
    <div>
      <div className={styles.toc}>
        <Toc />
      </div>
      <div className={styles.contentContainer}>{children}</div>
    </div>
  );
};

export default DocumentationLayout;
