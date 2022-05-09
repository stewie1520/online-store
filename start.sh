if [ -z "$ONLINE_STORE_DEBUG_MODE"  ]; then
  make
else
  echo "start application in debug mode"
  reflex -c ./reflex.conf
fi
