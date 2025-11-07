SELECT
    u.id,
    u.user_name,
    u.email,
    o.phone_number,
    o.road_number,
    o.holding_number,
    o.area,
    o.thana,
    o.district,
    b.title,
    b.author,
    b.description,
    b.image_path,
    b.category,
    b.is_stock,
    oi.quantity,
    oi.unit_price,
    oi.total_price,
    o.created_at
FROM
    users u
    JOIN orders o ON o.user_id = u.id
    JOIN order_items oi ON oi.order_id = o.id
    JOIN books b ON b.id = oi.book_id
ORDER BY
    u.id;