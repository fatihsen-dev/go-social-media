import React from "react";
import NextLogo from "./NextLogo";
import Link from "next/link";
import styles from "./Navbar.module.scss";

export default function Navbar() {
   return (
      <nav className={styles.Navbar}>
         <NextLogo className={styles} />
         <ul className={styles.Menu}>
            <li>
               <Link href='/explore'>Explore</Link>
            </li>
            <li>
               <Link href='/friends'>Friends</Link>
            </li>
            <li>
               <Link href='/profile'>Profile</Link>
            </li>
         </ul>
      </nav>
   );
}
