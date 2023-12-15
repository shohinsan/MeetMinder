import type { NavItem } from '$customTypes';

interface NavigationConfig {
    headerNavigation: NavItem[];
    sidebarTopNavigation: NavItem[];
    sidebarNavigation: Record<string, NavItem[]>; 
    sidebarBottomNavigation: NavItem[];
    dropdownNavigation: NavItem[];
    dropdownLogout: NavItem[];
    mobileNavigation: NavItem[];
}

export const config: NavigationConfig = {
    // Landing
    headerNavigation: [
        {
			name: 'Home',
			href: '/'
		},
		{
			name: 'About',
			href: '/about'
		},
		{
			name: 'Blogs',
			href: '/blog'
		},
		{
			name: 'Pricing',
			href: '/pricing'
		}
    ],
    // Dashboard - Desktop Sidebar
    sidebarTopNavigation: [
        { name: 'Events', href: 'events' },
        { name: 'Scheduled', href: 'scheduled' },
        { name: 'Availability', href: 'availability' }
    ],
    sidebarNavigation: {
        'This team': [
            { name: 'Heroicons', href: '#' },
            { name: 'Tailwind Labs', href: '#' },
            { name: 'Workcation', href: '#' }
        ],
    }, 
    sidebarBottomNavigation: [
        { name: 'Settings', href: 'settings' },
    ],
    // Dashboard - Dropdown
    dropdownNavigation: [
        { name: 'Account settings', href: 'settings' },
        { name: 'Support', href: '#' },
        { name: 'License', href: '#' },
    ],
    dropdownLogout: [
        { name: 'Logout', href: '#' },
    ],
    // Dashboard - Mobile Bottom Navigation
    mobileNavigation: [
        { name: 'Events', href: 'events' },
        { name: 'Scheduled', href: 'scheduled' },
        { name: 'Availability', href: 'availability' },
        { name: 'Settings', href: 'settings' },
    ]
};
