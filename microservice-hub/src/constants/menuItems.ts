export type MenuItem = {
  title: string
  subtitle: string
  icon: string
  route: string
}

export const menuItems: MenuItem[] = [
  {
    title: 'Orders',
    subtitle: 'View and manage orders',
    icon: 'mdi-cart-outline',
    route: '/orders',
  },
  {
    title: 'Items',
    subtitle: 'Browse available items',
    icon: 'mdi-package-variant',
    route: '/items',
  },
  {
    title: 'Inventory',
    subtitle: 'Track stock and inventory',
    icon: 'mdi-warehouse',
    route: '/inventory',
  },
  {
    title: 'Users',
    subtitle: 'Manage user accounts',
    icon: 'mdi-account-multiple-outline',
    route: '/users',
  },
]
