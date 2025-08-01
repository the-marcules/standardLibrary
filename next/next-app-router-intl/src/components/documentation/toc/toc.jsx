import React from 'react';

const tocLinks = [
    { href: '#introduction', label: 'Introduction' },
    { href: '#usage', label: 'Usage' },
    { href: '#examples', label: 'Examples' },
    { href: '#api', label: 'API' },
];

const Toc = () => (
    <nav>
        <ul>
            {tocLinks.map(link => (
                <li key={link.href}>
                    <a href={link.href}>{link.label}</a>
                </li>
            ))}
        </ul>
    </nav>
);

export default Toc;